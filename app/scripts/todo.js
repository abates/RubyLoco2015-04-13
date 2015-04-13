angular.module('todoApp', [
    'ngResource'
  ])
  .factory('Todo', ['$resource', function($resource) {
    return $resource('/todo/:id', { id: '@Id' } )
  }])
  .controller('TodoController', ['$scope', 'Todo', function($scope, Todo) {
    $scope.todoList = {};

    function resetNewTodo() {
      $scope.newTodo = {
        Task: ''
      }
    };

    $scope.refresh = function() {
      $scope.todoList = Todo.query();
    };

    $scope.add = function() {
      if ($scope.newTodo.Task.length > 0 && $scope.newTodo.Task.replace(/\s/g, '').length > 0) {
        Todo.save($scope.newTodo, function() {
          $scope.refresh();
        });
      }
      resetNewTodo();
    };

    $scope.delete = function(todo) {
      todo.$delete(function() {
        $scope.refresh();
      });
    };


    resetNewTodo();
    $scope.refresh();
  }]);
