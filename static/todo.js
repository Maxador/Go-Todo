function TaskController($scope, $http) {
	$scope.tasks = [];
	$scope.working = false;

	var logError = function(data, status) {
		console.log('Code ' + status + ': ' + data);
		$scope.working = false;
	};

	var refresh = function() {
		return $http.get('/task/').
			success(function(data) { $scope.tasks = data.Tasks; }).
			error(logError);
	};

	$scope.addTodo = function() {
		$scope.working = true;
		$http.post('/task/', {Title: $scope.taskTitle}).
			error(logError).
			success(function() {
				refresh().then(function() {
					$scope.working = false;
					$scope.taskTitle = '';
				});
			});
	};

	$scope.toggleDone = function(task) {
		data = {ID: task.ID, Title: task.Title, Done: !task.Done};
		$http.put('/task/' + task.ID, data).
			error(logError).
			success(function() {task.Done = !task.Done});		
	};

	refresh().then(function() {$scope.working = false; });
}