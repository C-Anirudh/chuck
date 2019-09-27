var app = angular.module('chuck', ['ngRoute']);

var global = {
    url: 'http://0.0.0.0:9000'
};

app.config(function($routeProvider) {
    $routeProvider
        .when("/", {
            templateUrl: 'html_components/home.html'
        })
        .when("/contact", {
            templateUrl: 'html_components/contact.html'
        })
        .when("/about", {
            templateUrl: 'html_components/about.html'
        })
        .when("/login", {
            templateUrl: 'html_components/user/login.html',
            controller: 'formController',
        })
        .when("/signup", {
            templateUrl: 'html_components/user/signup.html',
            controller: 'formController',
        })
        .when("/dashboard", {
            templateUrl: 'html_components/user/dashboard.html',
            controller: '',
        })
        .otherwise({
            templateUrl: 'html_components/error404.html'
        })
});

app.controller('formController', function($scope, $http, $location) {
    console.warn("Form Controller called.");

    $scope.handleLogin = function() {
        let data = 'email=' + $scope.user.email + '&password=' + $scope.user.password;
        console.log('data is', data);
        $http({
            url: global.url + '/login',
            method: 'POST',
            headers: {
                "Content-Type": "application/x-www-form-urlencoded"
            },
            data: data
        }).then(resp => {
            let res = resp.data;
            console.log('res is ', res)
            if (res) {
                $location.path('/dashboard');
            } else {
                $scope.error = 'Unable to log in';
            }
        });
    }

    $scope.handleSignup = function() {
        let data = 'name=' + $scope.newuser.name + '&email=' + $scope.newuser.email + '&password=' + $scope.newuser.password;
        console.log('data is', data);
        $http({
            url: global.url + '/signup',
            method: 'POST',
            headers: {
                'Content-Type': "application/x-www-form-urlencoded"
            },
            data: data
        }).then(resp => {
            let res = resp.data;
            console.log('res is', res)
            if (res) {
                $location.path('/login');
            } else {
                $scope.error = 'Unable to Sign Up';
            }
        });
    }
});