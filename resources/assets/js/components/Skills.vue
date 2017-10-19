<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Skills</h1>
            <p>
                My development skills on a scale of 1-5 stars. Note that the scale is relative to the other languages that I have used, so a 5 star rating means that I am most familiar with that language (Not that I am “perfect” at it)!
            </p>
            <span v-for="(skills, category) in skills">
                <hr>
                <h3>{{ category }}</h3>
                <div class="row">
                    <div class="col-sm-8 col-lg-4 mx-auto" v-for="(list, key) in skills" :class="[ key == 0 ? 'ml-lg-auto mr-lg-0' : 'mr-lg-auto ml-lg-0']">
                        <ul class="list-group list-group-inverse talent-list">
                            <li v-for="skill in list" class="list-group-item">
                                <span>{{ skill.title }}</span>
                                <div class="sr-only">{{ skill.rating }} of 5</div>
                                <div class="ability-score" aria-hidden="true" :title="skill.rating + ' Stars'">
                                    <i class="fa-star" aria-hidden="true" v-for="n in 5" :class="[ n <= skill.rating ? 'fas' : 'far empty']"></i>
                                </div>
                            </li>
                        </ul>
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
            skills: {},
        };
    },
    created: function() {
        this.$Progress.start()

        var vue = this
        axios.get('/api/skills')
            .then(function (response) {
                vue.$Progress.finish()
                vue.skills = response.data;
            })
            .catch(function (response) {
                vue.$Progress.fail()
            })
    },
    mounted: function() {
    }
}
</script>
