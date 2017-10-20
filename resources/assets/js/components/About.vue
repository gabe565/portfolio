<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1>About</h1>
            <hr>
            <div class="row">
                <div class="col-3 col-lg-2 ml-lg-auto mx-auto">
                    <img class="img-fluid rounded" src="/images/me.jpg">
                </div>
                <div class="col-lg-6 mr-auto">
                    <p>I am a <span>{{ age }}</span> year old computer programmer from Oklahoma City, OK. I am currently studying towards a B.S. in Computer Science at the University of Central Oklahoma.</p>
                    <p>I enjoy creating optimized programs which have an easy-to-use interface.</p>
                </div>
            </div>
            <hr>
            <div id="map">
                <gmap-map ref="map" :center="mapCenter" :options="mapOptions" style="width: 100%; height: 100%"></gmap-map>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    data: function() {
        return {
            age : '',
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
                ],
            }
        }
    },
    methods: {
        centerMap: _.throttle(function() {
            this.$refs.map.$mapObject.setCenter(this.mapCenter)
        }, 100, { leading: false })
    },
    created: function() {
        var birthday = +new Date('1995-05-26');
        this.age = ~~((Date.now() - birthday) / (31557600000));
    },
    mounted: function() {
        var vue = this
        $(window).resize(this.centerMap)
    },
    destroyed: function() {
        $(window).off('resize', this.centerMap)
    }
}
</script>
