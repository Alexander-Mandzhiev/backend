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
	tail = `FROM REGSKLMOV
        JOIN TARA on REGSKLMOV.REGSKLMOV_TARA = TARA.TARA_ID
        JOIN PART on PART.PART_ID = REGSKLMOV.REGSKLMOV_PART
        JOIN TOV on TOV.TOV_ID = REGSKLMOV.REGSKLMOV_TOV
        JOIN SKL on SKL.SKL_ID = REGSKLMOV.REGSKLMOV_SKL
        JOIN REGSKLOST on REGSKLOST.REGSKLOST_PART = REGSKLMOV.REGSKLMOV_PART
    WHERE PART.PART_START BETWEEN ? AND ?
        AND REGSKLMOV.REGSKLMOV_SKL = ?
        AND TARA.TARA_TYPE = 2809
        AND REGSKLMOV.REGSKLMOV_TYPE = 1`

	searchCondition = "AND (LOWER(part.part_name) LIKE LOWER(?) OR LOWER(tov.tov_name) LIKE LOWER(?))"
	groupByClause   = "GROUP BY tov.tov_name, part.part_name, PART.PART_START"
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

	query := `SELECT 
        MAX(PART.PART_START) AS max_part_start,
        SUM(REGSKLMOV.REGSKLMOV_KOL) AS total_kol,
        SUM(REGSKLMOV.REGSKLMOV_NET) AS total_net,
        tov.tov_name, 
        part.part_name,
        COUNT(*) AS count_tara 
    ` + tail

	if search != "" {
		query += " " + searchCondition
	}

	query += " " + groupByClause + " " + orderClause

	finalQuery := fmt.Sprintf(`
        SELECT FIRST %d SKIP %d 
            max_part_start,
            total_kol,
            total_net,
            tov_name,
            part_name,
            count_tara 
        FROM (%s)`, count, count*(page-1), query)

	args := []interface{}{dateStart, dateEnd, sklID}
	if search != "" {
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	rows, err := r.db.QueryContext(ctx, finalQuery, args...)
	if err != nil {
		sl.Log.Error(op, "failed to execute task query: ", err)
		return nil, fmt.Errorf("failed to execute task query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task production_task.Product
		var manufacturingDate time.Time
		var countTara int

		err = rows.Scan(
			&manufacturingDate,
			&task.CountSausageSticks,
			&task.WeightSpKg,
			&task.Nomenclature,
			&task.PartName,
			&countTara,
		)
		if err != nil {
			sl.Log.Error(op, "failed to scan task row: ", err)
			continue
		}

		task.ManufacturingDate = timestamppb.New(manufacturingDate)
		task.NumberFrame = fmt.Sprintf("%d", countTara) // Конвертация в строку
		tasks = append(tasks, &task)
	}

	if err = rows.Err(); err != nil {
		sl.Log.Error(op, "error during row iteration: ", err)
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
