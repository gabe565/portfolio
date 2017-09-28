(function($) {
    "use strict" // Start of use strict

    // Closes responsive menu when a scroll trigger link is clicked
    var navMain = $(".navbar-collapse") // avoid dependency on #id
    // "a:not([data-toggle])" - to avoid issues caused
    // when you have dropdown inside navbar
    navMain.on("click", "a:not([data-toggle])", null, function () {
        navMain.collapse('hide')
    })

    // Collapse the navbar when page is scrolled
    $(window).scroll(function() {
        if ($("#mainNav").offset().top > 100) {
            $("#mainNav").addClass("navbar-shrink")
        } else {
            $("#mainNav").removeClass("navbar-shrink")
        }
    })

})(jQuery) // End of use strict
