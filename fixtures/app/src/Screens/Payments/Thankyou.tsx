import { useEffect } from 'react'

export default function Thankyou() {
  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Thankyou',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Thankyou</h3>
    </div>
  )
}
