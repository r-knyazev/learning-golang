package requestService

type requestService struct {}

func newRequestService() RequestServiceInterface {
	return &requestService{}
}