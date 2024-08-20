package controllers

import (
	"myapp/models"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)



func FindPairAPi(c *gin.Context) {
	var request models.RequestData
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(request.Numbers)< 2{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Numbers Should be more than equal 2"})
		return
	}
	result := FindPair(request.Numbers, request.Target)
	c.JSON(http.StatusOK, gin.H{"solutions":result})
}

//returns all pairs of indiceswhere the sum of the elements at those indices equals the target value.
func FindPair(nums []int,target int) ( [][]int){
	var result = [][]int{}
	Mapindics := make(map[int]int) //store's the value and index
	for i, val := range nums {
		rem := target - val
		if j, ok := Mapindics[rem]; ok {
			result = append(result, []int{j, i})
		}
		Mapindics[val] = i
	}
	sort.Slice(result, func(i,j int) bool {
		return result[i][0] < result[j][0]
	})
	return  result
}