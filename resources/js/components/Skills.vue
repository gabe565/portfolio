<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Skills</h1>
            <div class="col-md-9 mx-auto">
                My development skills on a scale of 1-5 stars. Note that the scale is relative to the other languages that I have used, so a 5 star rating means that I am most familiar with that language (Not that I am perfect at it)!
            </div>
            <transition-group name="fade" v-if="ready" appear>
                <div v-for="(skills, category) in skillsChunked" :key="`cat-${category}`">
                    <hr>
                    <h3>{{ category }}</h3>
                    <div class="row">
                        <div class="col-sm-8 col-lg-4 mx-auto" v-for="(list, key) in skills" :class="[ key == 0 ? 'ml-lg-auto mr-lg-0' : 'mr-lg-auto ml-lg-0']">
                            <ul class="list-group list-group-inverse talent-list">
                                <li v-for="skill in list" class="list-group-item">
                                    <span>{{ skill.title }}</span>
                                    <div class="sr-only">{{ skill.rating }} of 5</div>
                                    <div class="ability-score" aria-hidden="true" :title="`${skill.rating} Stars`">
                                        <span v-for="n in 5">
                                            <template v-if="n <= skill.rating">
                                                <font-awesome-icon icon="fas fa-star" fixed-width class="text-primary"/>
                                            </template>
                                            <template v-else>
                                                <font-awesome-icon icon="far fa-star" fixed-width class="text-white opacity-15"/>
                                            </template>
                                        </span>
                                    </div>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>
            </transition-group>
        </div>
    </section>
</template>

<script>
export default {
    data() {
        return {
            ready: false,
            skills: null,
        }
    },
    computed: {
        skillsChunked() {
            if (this.ready)
                return _.mapValues(this.skills, (skills) => {
                    let pivot = Math.ceil(skills.length / 2)
                    return _.chunk(skills, pivot)
                })
            else
                return {}
        },
    },
    created() {
        this.$Progress.start()

        axios.get('/api/skills')
            .then(response => {
                this.skills = response.data
                this.ready = true
                this.$Progress.finish()
            })
            .catch(response => {
                this.$Progress.fail()
            })
    }
}
</script>

<style lang="scss">
.opacity-15 {
    opacity: 0.15;
}
</style>
