<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Projects</h1>
            <p>Below are links to the personal or collaborative projects that I can claim a piece of.</p>
            <span v-for="(projects, category) in projects">
                <hr>
                <h3>{{ category }}</h3>
                <transition-group name="fade" tag="div" appear class="card-deck col-lg-10 mx-auto">
                    <div class="card text-white bg-dark border-dark" v-for="project in projects" :key="`project-${project.name}`">
                        <a :href="project.url" class="card-link overflow-hidden" target="_blank">
                            <img :src="project.image_path" class="card-img-top" :alt="`Screenshot of ${project.name}`">
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
                </transition-group>
            </span>
        </div>
    </section>
</template>

<script>
import preloadImage from '../preloadImage'

export default {
    data() {
        return {
            projects: {},
        }
    },
    created() {
        this.$Progress.start()

        axios.get('/api/projects')
            .then(response => {
                _.forOwn(response.data, (category, key) => {
                    this.$set(this.projects, key, [])
                    _.forEach(category, project => {
                        preloadImage(project.image_path)
                            .then(() => {
                                this.projects[key].push(project)
                            })
                            .catch(() => {
                                this.$Progress.fail()
                            })

                    })
                })
                this.$Progress.finish()
            })
            .catch(response => {
                this.$Progress.fail()
            })
    }
}
</script>
