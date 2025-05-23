'use client'
import { JotCollection } from "@/types/jotTypes";
import { useState } from "react";

interface Props {
  jotCollection: JotCollection;
}

export default function JotDisplay({ jotCollection }: Props) {
  const [jots, _] = useState<JotCollection>(jotCollection)

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
              />
            ))}
          </div>
        </div>
      ))}
    </section>
  );
}
