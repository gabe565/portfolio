<template>
    <header class="masthead">
        <div class="intro-bg" :style="style"></div>
        <div class="intro-body">
            <div class="container">
                <div class="row">
                    <div class="col-lg-8 mx-auto">
                        <h1 class="brand-heading">Gabe Cook</h1>
                        <h2 class="intro-text">Software Developer</h2>
                    </div>
                </div>
            </div>
        </div>
    </header>
</template>

<script>
export default {
    data: function() {
        return {
            bgs: [
                'anas-alshanti-169265.jpg',
                'technology-785742.jpg'
            ]
        };
    },
    computed: {
        style: function() {
            return {
                'background-image':  'url(images/bg/' + _.sample(this.bgs) + ')',
                'opacity': 0.5
            }
        }
    }
}
</script>
