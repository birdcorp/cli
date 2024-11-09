var uniqueId = 1

export function log(
  message: any,
  data: any | undefined,
  type: any | undefined,
) {
  var log = document.getElementById('log')!
  var el = document.createElement('div')
  el.style.color = 'white'
  el.className = type === 'native' ? 'logLine_Native' : 'logLine_JS'
  el.innerHTML = uniqueId++ + '. ' + message + ':<br/>' + JSON.stringify(data)
  log.appendChild(el)
}
