package modules

import (
	"github.com/gin-gonic/gin"
)

// Module module interface
type Module interface {
	RegisterHandlers(r *gin.Engine)
}
