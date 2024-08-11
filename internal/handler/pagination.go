package handler

type Pagination struct {
	TotalPages  int
	CurrentPage int
	Pages       []int
}

func calculatePagination(totalItems, limit, offset int) Pagination {
	if totalItems == 0 {
		return Pagination{}
	}

	totalPages := (totalItems + limit - 1) / limit
	currentPage := (offset / limit) + 1

	pages := make([]int, totalPages)
	for i := 1; i <= totalPages; i++ {
		pages[i-1] = i
	}

	return Pagination{
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		Pages:       pages,
	}
}

func onPageClick(page int, limit int) int {
	return (page - 1) * limit
}
