<template>
    <div id="map">
        <gmap-map ref="map" :center="mapCenter" :options="mapOptions" style="width: 100%; height: 100%"></gmap-map>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            mapCenter: { lat:35.46756, lng:-97.516428 },
            mapOptions: {
                zoom: 11,
                disableDefaultUI: true,
                scrollwheel: false,
                draggable: false,
                disableDoubleClickZoom: true,
                styles: [
                    {
                        "featureType":"all",
                        "elementType":"all",
                        "stylers":[
                            { "invert_lightness":true },
                            { "saturation":10 },
                            { "lightness":30 },
                            { "gamma":0.5 },
                            { "hue":"#435158" }
                        ]
                    }
                ]
            }
        }
    },
    methods: {
        centerMap: _.throttle(function() {
            this.$refs.map.$mapObject.setCenter(this.mapCenter)
        }, 100, { leading: false })
    },
    mounted: function() {
        $(window).resize(this.centerMap)
    },
    destroyed: function() {
        $(window).off('resize', this.centerMap)
    }
}
</script>
