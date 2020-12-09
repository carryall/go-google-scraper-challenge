module.exports = function(grunt) {

  grunt.initConfig({
    sass: {
      dist: {
        options: {
          style: 'expanded'
        },
        files: {
          'static/css/application.css': 'assets/stylesheets/index.scss'
        },
      },
    },
    // PostCSS - Tailwindcss and Autoprefixer
    postcss: {
      options: {
        map: true, // inline sourcemaps
        processors: [
          require('tailwindcss')(),
          require('autoprefixer')({overrideBrowserslist: ['last 2 versions']}) // add vendor prefixes
        ]
      },
      dist: {
        files: {
          'static/css/tailwind.css': 'assets/stylesheets/vendors/tailwind.css'
        }
      }
    },
    watch: {
      scripts: {
        files: ['assets/stylesheets/**/**.scss'],
        tasks: ['sass', 'postcss'],
        options: {
          interrupt: true,
        },
      },
    },
  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-postcss');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['sass', 'postcss']);
};
