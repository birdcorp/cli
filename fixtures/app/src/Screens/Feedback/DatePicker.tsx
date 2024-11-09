import React, { useEffect, useCallback, useState } from 'react'
import { HighlightedCode } from '../../Components'
import { format } from 'date-fns'

export default function DatePicker() {
  const [selectedOption, setSelectedOption] = useState()

  useEffect(() => {
    window.miniProgram?.call('setNavigationBar', {
      title: 'Date Picker',
      backgroundColor: '#0A0A0A',
      color: 'light',
    })
  }, [])

  return (
    <div className="max-w-screen-sm mx-auto max-w-lg px-3 pt-5 bg-black text-gray-100">
      <h3 className="text-white font-bold mb-10">Date Picker</h3>

      {selectedOption && (
        <p className="text-white mb-10 text-xl">
          {format(new Date(selectedOption), 'MMMM do, yyyy h:mm a')}
        </p>
      )}
      <div className="flex flex-row items-center gap-2">
        <button
          onClick={() => {
            window.miniProgram?.call(
              'showDatePicker',
              {
                tintColor: '#FF5733', // Optional: HEX string for tint color
                // minDate: new Date('2023-01-01').getTime() / 1000, // Optional: Minimum selectable date (Unix timestamp)
                // maxDate: new Date('2024-12-31').getTime() / 1000, // Optional: Maximum selectable date (Unix timestamp)
                showTime: false, // Optional: Whether to show time picker (true/false)
              },
              {
                success: (response) => {
                  if (response.ok) {
                    setSelectedOption(response.result)
                  }
                },
              },
            )
          }}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Date
        </button>

        <button
          onClick={() => {
            window.miniProgram?.call(
              'showDatePicker',
              {
                tintColor: '#FF5733', // Optional: HEX string for tint color
                // minDate: new Date('2023-01-01').getTime() / 1000, // Optional: Minimum selectable date (Unix timestamp)
                // maxDate: new Date('2024-12-31').getTime() / 1000, // Optional: Maximum selectable date (Unix timestamp)
                showTime: true, // Optional: Whether to show time picker (true/false)
              },
              {
                success: (response) => {
                  if (response.ok) {
                    setSelectedOption(response.result)
                  }
                },
              },
            )
          }}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Date and Time
        </button>

        <button
          onClick={() => {
            window.miniProgram?.call(
              'showDatePicker',
              {
                tintColor: '#FF5733', // Optional: HEX string for tint color
                minDate: new Date('2023-01-01').getTime() / 1000, // Optional: Minimum selectable date (Unix timestamp)
                maxDate: new Date('2024-12-31').getTime() / 1000, // Optional: Maximum selectable date (Unix timestamp)
                showTime: false, // Optional: Whether to show time picker (true/false)
              },
              {
                success: (response) => {
                  if (response.ok) {
                    setSelectedOption(response.result)
                  }
                },
              },
            )
          }}
          className="bg-white text-neutral-800 font-bold py-2 px-4 rounded"
        >
          Date range
        </button>
      </div>

      <h3 className="text-white mt-8 font-bold">Code</h3>
      <HighlightedCode
        code={`window.miniProgram?.call(
'showDatePicker',
{
  tintColor: '#FF5733', // Optional: HEX string for tint color
  minDate: new Date('2023-01-01').getTime() / 1000, // Optional: Minimum selectable date (Unix timestamp)
  maxDate: new Date('2024-12-31').getTime() / 1000, // Optional: Maximum selectable date (Unix timestamp)
  showTime: true, // Optional: Whether to show time picker (true/false)
},
{
  success: (response) => {
    if (response.ok) {
      setSelectedOption(response.result)
    }
  },
},
)`}
      />
    </div>
  )
}
