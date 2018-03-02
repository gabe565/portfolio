<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1>About</h1>
            <hr>
            <div class="row">
                <div class="col-3 col-lg-2 mb-3 mx-auto ml-lg-auto mr-lg-0">
                    <img class="img-fluid rounded" src="/images/me.jpg">
                </div>
                <div class="col-lg-6 mr-auto">
                    <p>I am a <span>{{ age }}</span> year old computer programmer from Oklahoma City, OK. I am currently studying towards a B.S. in Computer Science at the University of Central Oklahoma.</p>
                    <p>I enjoy creating optimized programs which have an easy-to-use interface.</p>
                </div>
            </div>
            <hr>
        </div>
        <div id="map">
            <gmap-map ref="map" :center="mapCenter" :options="mapOptions" style="width: 100%; height: 100%"></gmap-map>
        </div>
    </section>
</template>

<script>
export default {
    data() {
        return {
            age : '',
            mapCenter: { lat: 35.46756, lng: -97.516428 },
            mapOptions: {
                zoom: 11,
                disableDefaultUI: true,
                scrollwheel: false,
                draggable: false,
                disableDoubleClickZoom: true,
                styles: [
                    {
                        "stylers": [
                            {
                                "hue": "#607d8b"
                            },
                            {
                                "saturation": 10
                            },
                            {
                                "invert_lightness": true
                            },
                            {
                                "lightness": 20
                            },
                            {
                                "gamma": 0.7
                            }
                        ]
                    },
                    {
                        "featureType": "administrative.land_parcel",
                        "elementType": "labels",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "poi",
                        "elementType": "labels.text",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "poi.business",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "road",
                        "elementType": "labels.icon",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "road.arterial",
                        "elementType": "labels",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "road.highway",
                        "elementType": "labels",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "road.local",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "road.local",
                        "elementType": "labels",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    },
                    {
                        "featureType": "transit",
                        "stylers": [
                            {
                                "visibility": "off"
                            }
                        ]
                    }
                ]}
        }
    },
    mounted() {
        var birthday = +new Date('1995-05-26')
        this.age = ~~((Date.now() - birthday) / (31557600000))
    }
}
</script>
