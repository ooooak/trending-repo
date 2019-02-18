var app = new Vue({
  el: '#app',
  data: {
    items: __output,
  },
  computed:{
    itemsKeys: function(){
      const date = function(date){
        const [d, m, y] = date.split("-")
        return new Date(y, m - 1, d);
      };
      return Object.keys(__output).sort((a, b) => {
        return date(a) > date(b) ? -1 : 1;
      });
    }, 
    items: function(){
      return __output
    }
  }
})