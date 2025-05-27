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
	baseConditions = `
  FROM REGSKLMOV
    JOIN TARA on REGSKLMOV.REGSKLMOV_TARA = TARA.TARA_ID
    JOIN PART on PART.PART_ID = REGSKLMOV.REGSKLMOV_PART
    JOIN TOV on TOV.TOV_ID = REGSKLMOV.REGSKLMOV_TOV
    JOIN SKL on SKL.SKL_ID = REGSKLMOV.REGSKLMOV_SKL
    JOIN REGSKLOST on REGSKLOST.REGSKLOST_PART = REGSKLMOV.REGSKLMOV_PART
  WHERE part.part_name = ?
    AND REGSKLMOV.REGSKLMOV_SKL = ?
    AND TARA.TARA_TYPE = 2809
    AND REGSKLMOV.REGSKLMOV_TYPE = 1
    AND PART.PART_START >= '01.01.2025'`

	orderClause = "ORDER BY TOV.TOV_NAME, PART.PART_NAME"
)

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
        PART.part_id,
        TOV.TOV_NAME,
        PART.PART_NAME,
        PART.PART_START,
        TARA.TARA_NAME,
        REGSKLMOV.REGSKLMOV_KOL,
        REGSKLMOV.REGSKLMOV_NET ` +
		baseConditions + " " + orderClause

	rows, err := r.db.QueryContext(ctx, query, params.Count, skip, params.PartName, params.GetSklId())
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var tasks []production_task.Product
	for rows.Next() {
		var task production_task.Product
		var manufacturingDate time.Time

		err = rows.Scan(&task.Id, &task.Nomenclature, &task.PartName, &manufacturingDate, &task.NumberFrame, &task.CountSausageSticks, &task.WeightSpKg)
		if err != nil {
			return nil, fmt.Errorf("row iteration error: %w", err)
		}

		task.ManufacturingDate = timestamppb.New(manufacturingDate)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *Repository) getTotalCount(ctx context.Context, params *production_task.RequestTaskParams) (int, error) {
	countQuery := `SELECT COUNT(*) ` + baseConditions

	var totalCount int
	countRow := r.db.QueryRowContext(ctx, countQuery, params.PartName, params.GetSklId())
	err := countRow.Scan(&totalCount)
	if err != nil {
		return 0, fmt.Errorf("failed to scan total count: %w", err)
	}
	return totalCount, nil
}
