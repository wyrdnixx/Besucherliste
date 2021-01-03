<template>
    <h2>Test</h2>
</template>

<script>
    import ManageCompany from './manageCompany.vue'

//import axios from 'axios';
    export default {
        name: 'index',
        components: {
            ManageCompany,
            ManageDepartments,
            ManageItems,
            ManageProcessors
        },
        data() {
            return {
                toggleViewCompany: false,
           
            }
        },
        methods: {
           
            }
        }
    }
</script>