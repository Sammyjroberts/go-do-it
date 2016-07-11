'use strict';

angular.module('app.todo').controller('todoCtrl',function($http, todoModel) {
  //init
  const self = this;
  self.showOptions = [{label:'Hide Completed',value:false},{label:'Show Completed',value:true}];
  self.showCompleted = self.showOptions[0];
  todoModel.getAll()
  .then( response => {
    console.log(response);
    self.todos = todoModel.formatTodos(response.data);
  });

  self.addNewTodo = function() {
    console.log('pushing');
    self.todos.push({
      name: '',
      due: new Date(),
      complete : false
    });
  };
  self.removeTodo = function(index) {
    console.log(index);
    _.pullAt(self.todos, index);
  };
});
