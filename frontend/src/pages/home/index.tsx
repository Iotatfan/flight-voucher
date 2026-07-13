import { type FormEvent, useState } from 'react'

interface ShortenerPageProps {
    apiBaseUrl: string
}

type Status = 'idle' | 'loading' | 'success' | 'already_exists' | 'error'

export function HomePage({ apiBaseUrl }: ShortenerPageProps) {
    const aircraftTypes = [
        "ATR",
        "Airbus 320",
        "Boeing 737 Max"
    ]

    const [crewName, setCrewName] = useState('')
    const [crewId, setCrewId] = useState('')
    const [flightNumber, setFlightNumber] = useState('')
    const [date, setDate] = useState('')
    const [aircraftType, setAircraftType] = useState('')
    const [status, setStatus] = useState<Status>('idle')
    const [errorMessage, setErrorMessage] = useState('')
    const [seats, setSeats] = useState<string[]>([])

    const getLocalDateString = () => {
        const today = new Date()
        const yyyy = today.getFullYear()
        const mm = String(today.getMonth() + 1).padStart(2, '0')
        const dd = String(today.getDate()).padStart(2, '0')
        return `${yyyy}-${mm}-${dd}`
    }
    const today = getLocalDateString()

    async function handleSubmit(e: FormEvent) {
        e.preventDefault()
        setStatus('loading')
        setErrorMessage('')
        setSeats([])

        if (date < today) {
            setErrorMessage('Flight date cannot be in the past')
            setStatus('error')
            return
        }

        try {
            // Step 1: check if flight already exists
            const checkRes = await fetch(`${apiBaseUrl}/api/check`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ flightNumber: flightNumber.trim(), date }),
            })
            if (!checkRes.ok) {
                const json = await checkRes.json().catch(() => ({}))
                throw new Error(json.message ?? `Server error ${checkRes.status}`)
            }
            const checkJson = await checkRes.json()

            if (checkJson.exists) {
                setStatus('already_exists')
                return
            }

            // Step 2: generate seats
            const genRes = await fetch(`${apiBaseUrl}/api/generate`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({
                    name: crewName.trim(),
                    id: crewId.trim(),
                    flightNumber: flightNumber.trim(),
                    date,
                    aircraft: aircraftType,
                }),
            })
            if (!genRes.ok) {
                const json = await genRes.json().catch(() => ({}))
                throw new Error(json.message ?? `Server error ${genRes.status}`)
            }
            const genJson = await genRes.json()
            setSeats(genJson.seats ?? [])
            setStatus('success')
        } catch (err: unknown) {
            setErrorMessage(err instanceof Error ? err.message : 'Something went wrong')
            setStatus('error')
        }
    }

    return (
        <main className="min-h-screen flex items-center justify-center bg-slate-50 px-4">
            <div className="w-full max-w-md bg-white rounded-2xl shadow-md border border-slate-200 p-8">
                <h1 className="text-2xl font-semibold text-slate-800 mb-1">Voucher Seat Assignent</h1>
                <p className="text-sm text-slate-500 mb-6">Generate Random Seats for Vouchers.</p>

                <form onSubmit={handleSubmit} className="flex flex-col gap-4">
                    <div className="flex flex-col gap-1">
                        <label htmlFor="flight-number" className="text-sm font-medium text-slate-700">
                            Crew Name
                        </label>
                        <input
                            id="crew-name"
                            type="text"
                            placeholder="e.g. John Smith"
                            value={crewName}
                            onChange={(e) => setCrewName(e.target.value)}
                            required
                            className="border border-slate-300 rounded-lg px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 transition"
                        />
                    </div>

                    <div className="flex flex-col gap-1">
                        <label htmlFor="flight-number" className="text-sm font-medium text-slate-700">
                            Crew ID
                        </label>
                        <input
                            id="crew-id"
                            type="text"
                            placeholder="e.g. 123456"
                            value={crewId}
                            onChange={(e) => setCrewId(e.target.value)}
                            required
                            className="border border-slate-300 rounded-lg px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 transition"
                        />
                    </div>

                    <div className="flex flex-col gap-1">
                        <label htmlFor="flight-number" className="text-sm font-medium text-slate-700">
                            Flight Number
                        </label>
                        <input
                            id="flight-number"
                            type="text"
                            placeholder="e.g. GA-123"
                            value={flightNumber}
                            onChange={(e) => setFlightNumber(e.target.value)}
                            required
                            className="border border-slate-300 rounded-lg px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 transition"
                        />
                    </div>

                    <div className="flex flex-col gap-1">
                        <label htmlFor="flight-date" className="text-sm font-medium text-slate-700">
                            Flight Date
                        </label>
                        <input
                            id="flight-date"
                            type="date"
                            value={date}
                            min={today}
                            onChange={(e) => setDate(e.target.value)}
                            required
                            className="border border-slate-300 rounded-lg px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 transition"
                        />
                    </div>

                    <div className="flex flex-col gap-1">
                        <label htmlFor="aircraft-type" className="text-sm font-medium text-slate-700">
                            Aircraft Type
                        </label>
                        <select
                            id="aircraft-type"
                            value={aircraftType}
                            onChange={(e) => setAircraftType(e.target.value)}
                            required
                            className="border border-slate-300 rounded-lg px-3 py-2 text-sm outline-none focus:ring-2 focus:ring-indigo-400 focus:border-indigo-400 transition bg-white"
                        >
                            <option value="" disabled>Select aircraft type</option>
                            {aircraftTypes.map((type) => (
                                <option key={type} value={type}>{type}</option>
                            ))}
                        </select>
                    </div>

                    <button
                        type="submit"
                        disabled={status === 'loading'}
                        className="bg-indigo-600 hover:bg-indigo-700 disabled:opacity-60 disabled:cursor-wait text-white font-medium text-sm rounded-lg px-4 py-2.5 transition"
                    >
                        {status === 'loading' ? 'Generating Voucher...' : 'Generate Voucher'}
                    </button>
                </form>

                {status === 'success' && seats.length > 0 && (
                    <div className="mt-5 rounded-lg bg-green-50 border border-green-200 px-4 py-3 text-sm text-green-700">
                        <p className="font-semibold mb-2">✓ Seats generated successfully</p>
                        <div className="flex gap-2 flex-wrap">
                            {seats.map((seat) => (
                                <span key={seat} className="inline-block bg-green-100 border border-green-300 text-green-800 font-mono font-medium rounded px-2 py-1">
                                    {seat}
                                </span>
                            ))}
                        </div>
                    </div>
                )}

                {status === 'already_exists' && (
                    <div className="mt-5 flex items-center gap-2 rounded-lg bg-red-50 border border-red-200 px-4 py-3 text-sm text-red-600">
                        <span>✕</span>
                        <span>Flight <strong>{flightNumber.toUpperCase()}</strong> on {date} already has a voucher assigned.</span>
                    </div>
                )}

                {status === 'error' && (
                    <div className="mt-5 rounded-lg bg-red-50 border border-red-200 px-4 py-3 text-sm text-red-600">
                        {errorMessage}
                    </div>
                )}
            </div>
        </main>
    )
}
