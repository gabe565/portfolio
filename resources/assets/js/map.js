// Google Maps Scripts
var map = null;
// When the window has finished loading create our google map below
google.maps.event.addDomListener(window, 'load', init);
google.maps.event.addDomListener(window, 'resize', function(){
    setTimeout(function() {
        map.setCenter(new google.maps.LatLng(35.46756,-97.516428));
    }, 0)
});

function init() {
    // Basic options for a simple Google Map
    // For more options see: https://developers.google.com/maps/documentation/javascript/reference#MapOptions
    var mapOptions = {
        // How zoomed in you want the map to start at (always required)
        zoom: 11,

        // The latitude and longitude to center the map (always required)
        center: new google.maps.LatLng(35.46756,-97.516428), // OKC

        // Disables the default Google Maps UI components
        disableDefaultUI: true,
        scrollwheel: false,
        draggable: false,
        disableDoubleClickZoom: true,
    };

    // Map style
    var styleOptions = {
        name: "Cobalt"
    };
    var mapStyle = [{"featureType":"all","elementType":"all","stylers":[{"invert_lightness":true},{"saturation":10},{"lightness":30},{"gamma":0.5},{"hue":"#435158"}]}];

    // Get the HTML DOM element that will contain your map
    // We are using a div with id="map" seen below in the <body>
    var $map = $('#map');

    // Create the Google Map using out element and options defined above
    map = new google.maps.Map($map, mapOptions);
    var mapType = new google.maps.StyledMapType(mapStyle, styleOptions);
    map.mapTypes.set("Cobalt", mapType);
    map.setMapTypeId("Cobalt");
};

