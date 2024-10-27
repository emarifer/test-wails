import { defineConfig } from 'vite'

export default defineConfig({
    build: {
        rollupOptions: {
            output: {
                entryFileNames: 'assets/htmx.min.js',
            }
        }
    }
})

// https://stackoverflow.com/questions/74760141/how-to-set-the-names-of-vite-build-js-and-css-bundle-names-in-vue
// https://v2.vitejs.dev/config/#config-intellisense
// https://stackoverflow.com/questions/70951072/vite-build-warns-script-cant-be-bundled-without-type-module-attribute