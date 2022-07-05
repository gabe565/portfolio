<template>
    <header class="masthead">
        <transition name="fade" appear>
            <div class="intro-bg" :key="index" :style="{ backgroundImage: `url(${backgrounds[index]})` }" v-if="index !== null"></div>
        </transition>
        <div class="intro-body">
            <div class="container">
                <div class="row">
                    <div class="col-lg-8 mx-auto">
                        <h1 class="brand-heading">{{ $parent.appName }}</h1>
                        <h2 class="intro-text">DevOps Engineer / Software Developer</h2>
                    </div>
                </div>
            </div>
        </div>
    </header>
</template>

<script>
import preloadImage from '../preloadImage'

export default {
    data() {
        return {
            index: null,
            backgrounds: null,
            timeout: null,
        }
    },
    created() {
        this.$Progress.start()
        this.updateBackgrounds()
    },
    methods: {
        async updateBackgrounds() {
            try {
                const response = await axios.get('/api/bg');
                this.backgrounds = response.data;
                await this.randomBackground();
            } catch (error) {
                console.error(error);
                this.$Progress.fail();
            }
        },
        async randomBackground() {
            let index = Math.floor(Math.random() * this.backgrounds.length)
            try {
                await preloadImage(this.backgrounds[index]);
                this.index = index;
                this.$Progress.finish();
                if (!this._inactive) {
                    this.startTimeout();
                }
            } catch (error) {
                console.error(error);
                this.$Progress.fail();
            }
        },
        startTimeout() {
            this.timeout = setTimeout(this.randomBackground, 7500)
        },
        stopTimeout() {
            if (this.timeout !== null)
                clearTimeout(this.timeout);
                this.timeout = null;
        }
    },
    activated() {
        if (this.index !== null && this.timeout === undefined) {
            this.startTimeout()
        }
    },
    deactivated() {
        this.stopTimeout()
        this.randomBackground()
    }
}
</script>
