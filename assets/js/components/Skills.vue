<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1 class="display-4">Skills</h1>
            <p>
                My development skills on a scale of 1-5 stars. Note that the scale is relative to the other languages that I have used, so a 5 star rating means that I am most familiar with that language (Not that I am “perfect” at it)!
            </p>
            <span v-for="(skills, category) in chunked">
                <hr>
                <h3>{{ category }}</h3>
                <div class="row">
                    <div class="col-sm-8 col-lg-4 mx-auto" v-for="(list, key) in skills" :class="[ key == 0 ? 'ml-lg-auto mr-lg-0' : 'mr-lg-auto ml-lg-0']">
                        <ul class="list-group list-group-inverse talent-list">
                            <li v-for="skill in list" class="list-group-item">
                                <span>{{ skill.title }}</span>
                                <div class="sr-only">{{ skill.rating }} of 5</div>
                                <div class="ability-score" aria-hidden="true" :title="skill.rating + ' Stars'">
                                    <span v-for="n in 5">
                                        <template v-if="n <= skill.rating">
                                            <svgicon name="star-full" class="star-full" aria-hidden="true"></svgicon>
                                        </template>
                                        <template v-else>
                                            <svgicon name="star-empty" class="star-empty" aria-hidden="true"></svgicon>
                                        </template>
                                    </span>
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
import '../svg/star-full'
import '../svg/star-empty'

export default {
    data: function() {
        return {
            skills: require('../json/skills.json'),
        }
    },
    computed: {
        chunked: function() {
            return _.mapValues(this.skills, function(value) {
                return _.chunk(value, _.ceil(_.size(value) / 2))
            })
        }
    }
}
</script>
