<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Projects</h1>
            <hr>
            <div class="card-deck">
                <div class="card text-white bg-dark border-dark" v-for="project in projects">
                    <a :href="project.url" class="card-link" target="_blank">
                        <img :src="project.image_path" class="card-img-top" :alt="'Screenshot of ' + project.name">
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
        </div>
    </section>
</template>

<script>
export default {
    data: function() {
        return {
            projects: {},
        };
    },
    created: function() {
        var vue = this
        axios.get('/api/projects')
            .then(function (response) {
                vue.projects = response.data;
            })
            .catch(function (response) {
            })
    }
}
</script>
