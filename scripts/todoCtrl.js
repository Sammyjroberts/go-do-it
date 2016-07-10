angular.module("app").controller("toDoCtrl",function($http) {
  const self = this;

  $http.get("/todos")
  .then(function(response){
    self.todos = response;
  })
})
