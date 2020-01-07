
// Hello Haskell
const concat = (arr1, arr2) => arr1.concat(arr2)
const flatten = obj => Object.values(obj).reduce(concat);
const sortCallback = (a, b) => a.stars > b.stars ? -1 : 1;
const sortStatsByTotal = (a, b) => b[1].total - a[1].total;
const sortStatsByTotalStars = (a, b) => b[1].totalStars - a[1].totalStars;

const val = function(value, _default){
  if (value == null || value == void 0){
    return _default;
  }
  return value;
}

// Sort dates
const sortDate = (a, b) => date(a) > date(b) ? -1 : 1;

const date = (date) => {
  const [d, m, y] = date.split("-")
  return new Date(y, m - 1, d);
};

const cleanEntries = function(storage){
  let ret = {};
  entryKeys.forEach(k => {
    ret[k] = val(storage[k], []).sort(sortCallback);
  });
  return ret;
};

const makeStats = function(ret, row){
  if (row != null && row.language != ""){
    let data = val(ret[row.language], {totalStars: 0, total: 0});      
    data.totalStars += row.stars;
    data.total += 1;
    ret[row.language] = data;
  }
  return ret;
}

const entryKeys = Object.keys(__output).sort(sortDate).splice(0, 3); 

var app = new Vue({
  el: '#app',
  data: {},
  computed:{
    itemsKeys: function(){
      return entryKeys;
    }, 
    items: function(){
      return cleanEntries(__output);
    },
    stats: function(){
      const coll = flatten(rows).reduce(makeStats, {});
      return Object.entries(coll).sort(sortStatsByTotalStars);
    }
  }
})