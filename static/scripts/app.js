'use strict';
angular.module('app',[
  'app.todo'
])
.constant('_', window._)
.run(function ($rootScope) {
   $rootScope._ = window._;
});
angular.module('app.todo',[]);
