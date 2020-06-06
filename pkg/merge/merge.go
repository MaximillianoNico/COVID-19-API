package group

import (
	"github.com/gin-gonic/gin"
)

type GroupGroup struct {
    groups []*gin.RouterGroup
}

func NewGroupGroup(groups []*gin.RouterGroup) GroupGroup {
    return GroupGroup {
        groups,
    }
}

func (g *GroupGroup) handle(method string, path string, handler gin.HandlerFunc) {
    for _, group := range g.groups {
        group.Handle(method, path, handler)
    }
}