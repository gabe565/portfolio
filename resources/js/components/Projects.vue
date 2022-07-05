<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Projects</h1>
            <p>Below are links to the personal or collaborative projects that I can claim a piece of.</p>
            <span v-for="(projects, category) in projects">
                <hr>
                <h3>{{ category }}</h3>

                <transition-group name="fade" tag="div" appear class="row justify-content-center">
                    <div class="col-10 col-lg-6 pb-4 pb-lg-0" v-for="project in projects" :key="`project-${project.name}`">
                        <div class="card text-white bg-dark border-dark h-100">
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
    async created() {
        this.$Progress.start()

        try {
            const response = await axios.get('/api/projects');
            _.forOwn(response.data, (category, key) => {
                this.$set(this.projects, key, []);
                _.forEach(category, async project => {
                    try {
                        await preloadImage(project.image_path);
                        this.projects[key].push(project);
                    } catch (error) {
                        console.error(error);
                        this.$Progress.fail();
                    }
                });
            });
            this.$Progress.finish();
        } catch (error) {
            console.error(error);
            this.$Progress.fail();
        }
    }
}
</script>
