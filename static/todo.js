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
		$score.working = true;
		$http.get('/task/' + task.ID)
			.success(function(data) {});
	};

	refresh().then(function() {$scope.working = false; });
}