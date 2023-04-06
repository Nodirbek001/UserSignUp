package api
import(
	"github.com/gorilla/mux"

)

type routes struct{
	root *mux.Router
	apiRoot *mux.Router
}
type api struct{
	routes *routes
	
}