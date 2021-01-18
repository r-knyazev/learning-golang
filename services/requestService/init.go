package requestService

var RequestService RequestServiceInterface

func init()  {
	RequestService = newRequestService()
}