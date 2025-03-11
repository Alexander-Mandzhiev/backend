package repository

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/production_task"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

const (
	tail = `FROM operfas
        LEFT JOIN tov ON tov.tov_id = operfas.operfas_srctov
        LEFT JOIN tara ON operfas.operfas_srctara = tara.tara_id
        LEFT JOIN part ON part.part_id = operfas.operfas_srcpart
        LEFT JOIN reports_recorded_mssql ON reports_recorded_mssql.operfas_id = operfas.operfas_id
    WHERE operfas.operfas_dt BETWEEN ? AND ?
        AND operfas.operfas_srcnet <> 0
        AND operfas.operfas_srcskl = ?
        AND tara.tara_type = 2809
        AND reports_recorded_mssql.operfas_id IS NULL`
	searchCondition = "LOWER(part.part_name) LIKE LOWER(?) OR LOWER(tov.tov_name) LIKE LOWER(?)"
	groupByClause   = "GROUP BY tov.tov_name, part.part_name"
)

func (r *Repository) Tasks(ctx context.Context, params *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	tasks, err := r.getTasks(ctx, params)
	if err != nil {
		return nil, err
	}

	totalCount, err := r.countTasks(ctx, params)
	if err != nil {
		return nil, err
	}

	totalPages := (totalCount + int(params.GetCount()) - 1) / int(params.GetCount())

	protoProducts := make([]*production_task.Product, 0, len(tasks))
	for _, task := range tasks {
		protoProducts = append(protoProducts, task)
	}

	return &production_task.ProductsResponse{
		Data:      protoProducts,
		TotalPage: int32(totalPages),
	}, nil
}

func (r *Repository) getTasks(ctx context.Context, params *production_task.RequestTaskParams) ([]*production_task.Product, error) {
	op := "FirebirdRepository.getTasks"
	var tasks []*production_task.Product

	page := int(params.GetPage())
	count := int(params.GetCount())
	sklID := params.GetSklId()
	search := params.GetSearch()
	dateStart := time.Unix(params.DateStart.Seconds, 0)
	dateEnd := time.Unix(params.DateEnd.Seconds, 0)

	skip := count * (page - 1)

	query := "SELECT FIRST ? SKIP ? MAX(operfas.operfas_dt) AS max_operfas_dt, SUM(operfas.operfas_srcnet) AS max_operfas_srcnet, tov.tov_name, part.part_name, COUNT(*) AS countTara " + tail

	var args []interface{}
	args = append(args, count, skip, dateStart, dateEnd, sklID)

	if search != "" {
		query += " AND (" + searchCondition + ")"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	query += " " + groupByClause

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		sl.Log.Error(op, "failed to execute task query: ", err)
		return nil, fmt.Errorf("failed to execute task query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task production_task.Product
		var manufacturingDate time.Time

		if err = rows.Scan(
			&manufacturingDate,
			&task.WeightSpKg,
			&task.Nomenclature,
			&task.PartName,
			&task.NumberFrame,
		); err != nil {
			sl.Log.Error(op, "failed to scan task row: ", err)
			continue
		}
		task.ManufacturingDate = timestamppb.New(manufacturingDate)
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error(op, "error encountered during row iteration: ", err)
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return tasks, nil
}

func (r *Repository) countTasks(ctx context.Context, params *production_task.RequestTaskParams) (int, error) {
	op := "FirebirdRepository.countTasks"
	var totalCount int

	sklID := params.GetSklId()
	search := params.GetSearch()
	dateStart := time.Unix(params.DateStart.Seconds, 0)
	dateEnd := time.Unix(params.DateEnd.Seconds, 0)

	query := "SELECT COUNT(*) AS total_count FROM (SELECT tov.tov_name, part.part_name " + tail

	var args []interface{}
	args = append(args, dateStart, dateEnd, sklID)

	if search != "" {
		query += " AND (" + searchCondition + ")"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	query += " " + groupByClause + ") AS grouped_data "

	countRow := r.db.QueryRowContext(ctx, query, args...)
	if err := countRow.Scan(&totalCount); err != nil {
		sl.Log.Error(op, "failed to retrieve total count: ", err)
		return 0, fmt.Errorf("failed to retrieve total count: %w", err)
	}

	return totalCount, nil
}
