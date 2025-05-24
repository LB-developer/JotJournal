'use client'
import { Jot, JotCollection } from "@/types/jotTypes";
import { useRouter } from "next/navigation";
import { useState } from "react";


interface Props {
  jotCollection: JotCollection;
}

export default function JotDisplay({ jotCollection }: Props) {
  const router = useRouter()
  const [jots, setJots] = useState<JotCollection>(jotCollection)

  const handleUpdateJot = async (
    e: React.MouseEvent<HTMLDivElement, MouseEvent>,
    jotToUpdate: Jot
  ): Promise<void> => {

    e.preventDefault()
    const jotID = jotToUpdate.id
    const habit = jotToUpdate.habit
    const changedCompletion = !jotToUpdate.isCompleted

    await fetch("/api/jots", 
      { 
        method: "PATCH",
        body: JSON.stringify({ jotID, isCompleted: changedCompletion })
      })

    router.refresh() // invalidate cache

    const updatedJots = {...jots}
    const length = updatedJots[habit].length - 1
    updateJot(updatedJots[habit], 0, length, jotID)
    setJots(updatedJots)
  }

  return (
    <section className="flex flex-row gap-4 p-4">
      {Object.entries(jots).map(([habit, jots]) => (
        <div key={habit}>
          <h4 className="text-sm font-semibold mb-1">{habit}</h4>
          <div className="gap-2">
            {jots.map((jot) => (
              <div
                key={jot.id}
                className={`w-5 h-5 rounded-sm border ${
                  jot.isCompleted
                    ? 'bg-green-500 border-green-700'
                    : 'bg-gray-200 border-gray-400'
                }`}
                title={jot.date}
                onClick={(e) => handleUpdateJot(e, jot)}
              />
            ))}
          </div>
        </div>
      ))}
    </section>
  );
}


function updateJot(haystack: Jot[], left:number, right:number, target: number) {
  let mid = left + Math.floor((right - left) / 2)

  if (haystack[mid].id === target) {
    // change the current jot
    haystack[mid].isCompleted = !haystack[mid].isCompleted
    return

  } else if (haystack[mid].id < target) {
    left = mid + 1
  } else if (haystack[mid].id > target) {
    right = mid
  }

  updateJot(haystack, left, right, target)
}
