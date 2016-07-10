angular.module("app").controller("toDoCtrl",function($http) {
  const self = this;

  $http.get("/todos")
  .then(function(response){
    console.log(response);
    response.data.forEach(function(resp) {
      resp.due = new Date(resp.due);
    })
    self.todos = response.data;

  })
})
