module.exports = {
    mode: 'jit',
    future: {
        removeDeprecatedGapUtilities: true,
        purgeLayersByDefault: true,
    },
    purge: [
        'views/**/*.tpl'
    ],
    theme: {
        extend: {},
    },
    variants: {},
    plugins: [],
}
