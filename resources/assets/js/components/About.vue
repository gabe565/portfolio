<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1 class="display-4">About</h1>
            <div class="row">
                <div class="col-3 col-lg-2 ml-lg-auto mx-auto">
                    <img class="img-fluid rounded" src="/images/me.jpg">
                </div>
                <div class="col mr-auto">
                    <p>I am a <span>{{ age }}</span> year old computer programmer from Oklahoma City, OK. I am currently studying towards a B.S. in Computer Science at the University of Central Oklahoma.</p>
                    <p>I enjoy creating optimized programs which have an easy-to-use interface.</p>
                </div>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    data: function() {
        return {
            age : '',
        }
    },
    mounted: function() {
        var birthday = +new Date('1995-05-26');
        this.age = ~~((Date.now() - birthday) / (31557600000));
    }
}
</script>
