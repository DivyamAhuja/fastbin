/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./internal/**/*.{go,js,templ,html}"
    ],
    theme: {
        extend: {
            colors: {
                "primary": "rgb(var(--primary-color) / <alpha-value>)",
                "border": "rgb(var(--border-color) / <alpha-value>)"
            }
        },
    },
    plugins: [],
}