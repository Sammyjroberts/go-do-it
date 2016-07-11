'use strict';

angular.module('app.todo').service('todoModel', function($http) {
  const self = this;
  const TODO_ROUTE = '/todos/';

  self.getOne = function() {
    return $http.get(TODO_ROUTE);
  };
  self.getAll = function() {
    return $http.get(TODO_ROUTE);
  };
  self.post = function(data) {
    return $http.post(TODO_ROUTE, data);
  };
  self.put = function(data, id) {
    return $http.put(TODO_ROUTE+id, data);
  };
  self.delete = function(data, id) {
    return $http.delete(TODO_ROUTE+id, data);
  };
  self.formatTodos = function(todos) {
    todos.forEach(function(todo) {
      todo.due = new Date(todo.due);
    });
    return todos;
  };

  return self;
});
