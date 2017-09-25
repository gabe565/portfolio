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
            style: {},
        };
    },
    created: function() {
        var vue = this
        axios.get('/api/bg')
            .then(function (response) {
                vue.style = {
                    'background-image':  'url(' + response.data + ')',
                    'opacity': 0.5
                };
            })
            .catch(function (response) {
            })
    }
}
</script>
