<template>
    <section id="skills" class="content-section text-center">
        <div class="container">
            <h1>Projects</h1>
            <span v-for="(skills, category) in skills">
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
                <hr>
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
        var vue = this
        axios.get('/api/skills')
            .then(function (response) {
                vue.skills = response.data;
            })
            .catch(function (response) {
            })
    }
}
</script>
