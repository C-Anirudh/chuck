var app = angular.module('chuck', ['ngRoute']);

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
            templateUrl: 'html_components/login.html'
        })
        .when("/signup", {
            templateUrl: 'html_components/signup.html'
        })
        .otherwise({
            templateUrl: 'html_components/error404.html'
        })
});

$(document).ready(function() {
    $.each($('#navbar').find('li'), function() {
        $(this).toggleClass('active', 
            window.location.pathname.indexOf($(this).find('a').attr('href')) > -1);
    }); 
});