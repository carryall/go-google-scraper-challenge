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
    watch: {
      scripts: {
        files: ['assets/stylesheets/**/**.scss'],
        tasks: ['sass'],
        options: {
          interrupt: true,
        },
      },
    },
  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['sass']);

};
