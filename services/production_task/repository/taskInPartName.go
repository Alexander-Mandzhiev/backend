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
	startDate = "2024-01-01"
)

func getCurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func (r *Repository) TaskInPartName(ctx context.Context, params *production_task.RequestTaskParams) (*production_task.ProductsResponse, error) {
	op := "FirebirdRepository.TaskInPartName"

	page := int(params.GetPage())
	count := int(params.GetCount())
	skip := count * (page - 1)

	tasks, err := r.getTasksInPName(ctx, params, skip)
	if err != nil {
		sl.Log.Error(op, "failed to get tasks: ", err.Error())
		return nil, err
	}

	totalCount, err := r.getTotalCount(ctx, params)
	if err != nil {
		sl.Log.Error(op, "failed to retrieve total count: ", err.Error())
		return nil, err
	}

	totalPages := (totalCount + count - 1) / count

	protoProducts := make([]*production_task.Product, 0, len(tasks))
	for _, task := range tasks {
		protoProducts = append(protoProducts, &task)
	}

	return &production_task.ProductsResponse{
		Data:      protoProducts,
		TotalPage: int32(totalPages),
	}, nil
}

func (r *Repository) getTasksInPName(ctx context.Context, params *production_task.RequestTaskParams, skip int) ([]production_task.Product, error) {
	query := `SELECT FIRST ? SKIP ? 
        operfas.operfas_id,
        MAX(operfas.operfas_dt) AS max_operfas_dt, 
        MAX(operfas.operfas_srcnet) AS max_operfas_srcnet,
        tov.tov_name, 
        part.part_name,
        tara.tara_code
    FROM operfas 
    LEFT JOIN tov ON tov.tov_id = operfas.operfas_srctov 
    LEFT JOIN tara ON operfas.operfas_srctara = tara.tara_id 
    LEFT JOIN part ON part.part_id = operfas.operfas_srcpart 
    LEFT JOIN reports_recorded_mssql ON reports_recorded_mssql.operfas_id = operfas.operfas_id 
    WHERE operfas.operfas_dt BETWEEN ? AND ?
        AND operfas.operfas_srcnet <> 0 
        AND part.part_name = ?
        AND operfas.operfas_srcskl = ?
        AND tara.tara_type = 2809
        AND reports_recorded_mssql.operfas_id IS NULL
    GROUP BY tov.tov_name, part.part_name, tara.tara_code, operfas.operfas_id;`

	rows, err := r.db.QueryContext(ctx, query, params.Count, skip, startDate, getCurrentDate(), params.PartName, params.GetSklId())
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var tasks []production_task.Product
	for rows.Next() {
		var task production_task.Product
		var manufacturingDate time.Time

		err = rows.Scan(&task.Id, &manufacturingDate, &task.WeightSpKg, &task.Nomenclature, &task.PartName, &task.NumberFrame)
		if err != nil {
			return nil, fmt.Errorf("row iteration error: %w", err)
		}

		task.ManufacturingDate = timestamppb.New(manufacturingDate)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *Repository) getTotalCount(ctx context.Context, params *production_task.RequestTaskParams) (int, error) {
	countQuery := `SELECT COUNT(*) 
        FROM operfas 
        LEFT JOIN tov ON tov.tov_id = operfas.operfas_srctov 
        LEFT JOIN tara ON operfas.operfas_srctara = tara.tara_id 
        LEFT JOIN part ON part.part_id = operfas.operfas_srcpart 
        LEFT JOIN reports_recorded_mssql ON reports_recorded_mssql.operfas_id = operfas.operfas_id 
        WHERE operfas.operfas_dt BETWEEN ? AND ?
            AND operfas.operfas_srcnet <> 0 
            AND part.part_name = ?
            AND operfas.operfas_srcskl = ?
            AND tara.tara_type = 2809
            AND reports_recorded_mssql.operfas_id IS NULL;`

	var totalCount int
	countRow := r.db.QueryRowContext(ctx, countQuery, startDate, getCurrentDate(), params.PartName, params.GetSklId())
	err := countRow.Scan(&totalCount)
	if err != nil {
		return 0, fmt.Errorf("failed to scan total count: %w", err)
	}
	return totalCount, nil
}
