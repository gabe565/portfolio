// Google Maps Scripts
let map = null
// When the window has finished loading create our google map below
//google.maps.event.addDomListener(window, 'load', init)
google.maps.event.addDomListener(window, 'resize', () => {
    setTimeout(() => {
        map.setCenter(new google.maps.LatLng(35.46756,-97.516428))
    }, 0)
})

//function init() {
    // Basic options for a simple Google Map
    // For more options see: https://developers.google.com/maps/documentation/javascript/reference#MapOptions
    let mapOptions = {
        // How zoomed in you want the map to start at (always required)
        zoom: 11,

        // The latitude and longitude to center the map (always required)
        center: new google.maps.LatLng(35.46756,-97.516428), // OKC

        // Disables the default Google Maps UI components
        disableDefaultUI: true,
        scrollwheel: false,
        draggable: false,
        disableDoubleClickZoom: true,
    }

    // Map style
    let styleOptions = {
        name: "Cobalt"
    }
    let mapStyle = [
        {
            "featureType": "all",
            "elementType": "all",
            "stylers": [
                {
                    "invert_lightness": true
                },
                {
                    "saturation": 10
                },
                {
                    "lightness": 30
                },
                {
                    "gamma": 0.5
                },
                {
                    "hue":"#435158"
                }
            ]
        }
    ]

    // Get the HTML DOM element that will contain your map
    // We are using a div with id="map" seen below in the <body>
    let $map = document.getElementById('map')

    // Create the Google Map using out element and options defined above
    window.map = new google.maps.Map($map, mapOptions)
    let mapType = new google.maps.StyledMapType(mapStyle, styleOptions)
    window.map.mapTypes.set("Cobalt", mapType)
    window.map.setMapTypeId("Cobalt")
//}
