package main

import (
	"fmt"
	"math"
)

type RequestPagination struct {
	Page   int64 `json:"page"`   // current page
	Size   int64 `json:"size"`   // limit
	Offset int64 `json:"offset"` // offset
	Total  int64 `json:"total"`  // total row
}

type Pagination struct {
	Page       int64 `json:"page"`       // current page
	Size       int64 `json:"size"`       // limit
	Offset     int64 `json:"offset"`     // offset
	TotalPages int64 `json:"totalPages"` // total page
	Total      int64 `json:"total"`      // total row
	Visible    int64 `json:"visible"`    // total row in current page
	Last       bool  `json:"last"`       // is last page
	First      bool  `json:"first"`      // is first page
}

func CreatePagination(req_pagination RequestPagination) (Pagination, error) {
	if req_pagination.Size <= req_pagination.Total {
		if req_pagination.Offset <= req_pagination.Total {
			var pagination Pagination
			if req_pagination.Page == 0 {
				pagination.Page = 1
			} else {
				pagination.Page = req_pagination.Page
			}
			pagination.Size = req_pagination.Size
			pagination.Offset = req_pagination.Offset
			pagination.Total = req_pagination.Total
			if pagination.Total <= pagination.Size {
				pagination.Visible = pagination.Total
			} else if pagination.Total > pagination.Size {
				current_total := pagination.Page * pagination.Size
				if pagination.Total > current_total {
					pagination.Visible = pagination.Size
				} else {
					mod_total := pagination.Total % pagination.Size
					pagination.Visible = mod_total
				}
			}
			total_pages := math.Ceil(float64(pagination.Total / pagination.Size))
			pagination.TotalPages = int64(total_pages)
			if pagination.Page == 1 {
				pagination.First = true
				pagination.Last = false
			} else if pagination.Page == pagination.TotalPages {
				pagination.First = false
				pagination.Last = true
			} else {
				pagination.First = false
				pagination.Last = false
			}
			return pagination, nil
		} else {
			return Pagination{}, fmt.Errorf("offset have to less than or equals to total")
		}
	} else {
		return Pagination{}, fmt.Errorf("size have to less than or equals to total")
	}

}
