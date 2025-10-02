package structs

// Struct ini digunakan untuk menampilkan data user sebagai response API
type PostResponse struct {
	Id        uint   `json:"id"`
	Title     string `json:"name"`
	Body      string `json:"username"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Struct ini digunakan untuk menerima data saat proses create user
type PostCreateRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// Struct ini digunakan untuk menerima data saat proses update user
type PostUpdateRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}
