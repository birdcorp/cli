import React, { useEffect } from 'react'
import { HighlightedCode } from '../../Components'

export default function Vibrate() {
  const code = `window.miniprogram?.call("vibrate", {
    type: 'success'
})`

  function trigger(
    type:
      | 'light'
      | 'medium'
      | 'heavy'
      | 'success'
      | 'warning'
      | 'error'
      | 'selection',
  ) {
    window.miniProgram?.call(
      'vibrate',
      {
        type,
      },
      {
        success(response: any) {
          if (response.ok) {
            // confirm button was pressed
          }
        },
        failure(err: any) {
          console.error(err)
        },
        completed() {},
      },
    )
  }

  useEffect(() => {
    async function onMount() {
      window.miniProgram?.call(
        'setNavigationBar',
        {
          title: 'Vibrate',
          backgroundColor: '#0A0A0A',
          color: 'light',
        },
        {
          success(response: any) {
            if (response.ok) {
              // confirm button was pressed
            }
          },
          failure(err: any) {
            console.error(err)
          },
          completed() {},
        },
      )
    }

    onMount()
  }, [])
  /*
 
 HapticFeedbackManager.shared.impact(style: .light)
 HapticFeedbackManager.shared.impact(style: .medium)
 HapticFeedbackManager.shared.impact(style: .heavy)
 
 HapticFeedbackManager.shared.notification(type: .success)
 HapticFeedbackManager.shared.notification(type: .warning)
 HapticFeedbackManager.shared.notification(type: .error)
 
 HapticFeedbackManager.shared.selection()
 
 */

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-3">Vibrate</h3>

      <p className="text-gray-300">Triggers haptic feedback vibration</p>

      <div>
        <table className="table-auto text-neutral-300 w-full mt-8">
          <thead>
            <tr>
              <th>Property</th>
              <th>Value</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>type</td>
              <td>light, medium, heavy, success, warning, error, selection</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div className="flex flex-row justify-between mt-8">
        <h3 className="text-white font-bold">Code</h3>
        <p className="text-gray-300">Javascript</p>
      </div>

      <HighlightedCode code={code} />

      <div className="flex flex-col mt-8">
        <button
          onClick={() => trigger('light')}
          className="bg-gray-600 hover:bg-gray-500 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Light
        </button>

        <button
          onClick={() => trigger('medium')}
          className="bg-gray-700 hover:bg-gray-600 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Medium
        </button>

        <button
          onClick={() => trigger('heavy')}
          className="bg-gray-800 hover:bg-gray-700 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Heavy
        </button>
      </div>

      <div className="flex flex-col mt-8">
        <button
          onClick={() => trigger('success')}
          className="bg-green-600 hover:bg-green-500 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Success
        </button>

        <button
          onClick={() => trigger('warning')}
          className="bg-orange-600 hover:bg-orange-500 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Warning
        </button>

        <button
          onClick={() => trigger('error')}
          className="bg-red-600 hover:bg-red-500 text-white font-bold py-4 rounded-3xl mb-2"
        >
          Error
        </button>
      </div>

      <div className="flex flex-col mt-8">
        <button
          onClick={() => trigger('selection')}
          className="bg-gray-800 hover:bg-gray-700 text-gray-300 font-bold py-4 rounded-3xl mb-2"
        >
          Selection
        </button>
      </div>
    </div>
  )
}
