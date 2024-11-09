import React, { useEffect, useState } from 'react'
import { HighlightedCode } from '../../Components'

import { ScanLine } from 'lucide-react'

export default function Scan() {
  const [result, setResult] = useState<string | null>(null)

  function openScanner() {
    window.miniProgram?.call(
      'scan',
      {
        type: 'qr',
      },
      {
        success(response: any) {
          if (response.ok) {
            setResult(response.result)
          }
        },
      },
    )
  }

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Scan',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 text-gray-100">
      <h3 className="text-white font-bold mb-3">Scan</h3>

      <p className="text-gray-400 mb-8">Triggers a scanner</p>

      {result && (
        <div className="flex flex-col items-start justify-start">
          <h3>Response:</h3>
          <p className="text-gray-400 mb-8 text-3xl">{result}</p>
        </div>
      )}

      <button
        onClick={openScanner}
        className="bg-white text-neutral-800 font-bold py-2 px-4 rounded flex flex-row items-center gap-2"
      >
        <ScanLine />
      </button>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-400">Javascript</p>
      </div>

      <HighlightedCode
        code={`window.miniProgram?.call('scan', 
      { type: 'qr' },
      {
        success(response) {
          if (response.ok) {
            // response.result
          }
        }
      },
    )`}
      />
    </div>
  )
}
