<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1 class="display-4">Projects</h1>
            <p>
                Below are links to the personal or collaborative projects that I can claim.
            </p>
            <span v-for="(projects, category) in projects">
                <hr>
                <h3>{{ category }}</h3>
                <div class="card-deck">
                    <div class="card text-white bg-dark border-dark" v-for="project in projects">
                        <a :href="project.url" class="card-link overflow-hidden" target="_blank" :style="{ height: cardHeight }">
                            <img :src="project.image_path" class="card-img-top" :alt="'Screenshot of ' + project.name" v-on:load="cardEqHeight">
                        </a>
                        <div class="card-body">
                            <h5 class="card-title">{{ project.name }}</h5>
                            <p class="card-description" v-html="project.description"></p>
                        </div>
                        <a :href="project.url" class="card-link" target="_blank">
                            <div class="card-footer">
                                    View at <span class="mono">{{ project.url.replace(/(^\w+:|^)\/\//, '') }}</span>
                            </div>
                        </a>
                    </div>
                </div>
            </span>
        </div>
    </section>
</template>

<script>
export default {
    data: function() {
        return {
            cardHeight: 0,
            projects: require('../json/projects.json'),
        };
    },
    methods: {
        cardEqHeight: function() {
            this.cardHeight = $('.card-img-top').first().width() / 1.5 + "px"
        }
    },
    mounted: function() {
        $(window).on('resize.projects', _.throttle(this.cardEqHeight, 50, { leading: false, trailing: true }))
    },
    destroyed: function() {
        $(window).off('resize.projects')
    }
}
</script>
