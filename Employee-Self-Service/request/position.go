package request

import "employeeSelfService/domain"

type Position struct {
	IdPosition   int64  `json:"id_position"`
	PositionName string `json:"position_name"`
}

func (ps Position) ToDomainPosition() *domain.Position {
	return &domain.Position{
		IdPosition:   ps.IdPosition,
		PositionName: ps.PositionName,
	}
}
