import { useFormStatus } from "react-dom"

interface Props {
  actionText: string
}

export function SubmitButton({actionText}: Props) {
    const { pending } = useFormStatus()

    return (
        <button 
          className="w-full bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2.5 rounded-lg transition-colors"
          disabled={pending} 
          type="submit"
        >
          {pending ? "Submitting" : actionText}
        </button>
    )
}
