module.exports = {
    content: [
        "./resources/templates/**/*.temple",
        "./resources/js/**/*.js",
    ],
    theme: {
        extend: {},
    },
    plugins: [],
    // Используйте purge, чтобы убрать неиспользуемые классы
    purge: ['./resources/templates/**/*.{gohtml,js}', './resources/js/**/*.{js}'],
}
