import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "@/components/ui/card"
import { Badge } from "@/components/ui/badge"
import { Separator } from "@/components/ui/separator"
import { Input } from "@/components/ui/input"
import { Select, SelectTrigger, SelectValue, SelectContent, SelectItem } from "@/components/ui/select"
import { Checkbox } from "@/components/ui/checkbox"
import { Textarea } from "@/components/ui/textarea"
import { FormField, FormItem, FormLabel, FormControl, FormMessage } from "@/components/ui/form"
import { Question } from "@/types/screening"

interface ScreeningQuestionsProps {
  questions: Question[]
  form: any
  step: number
}

export function ScreeningQuestions({ questions, form, step }: ScreeningQuestionsProps) {
  if (step !== 2 || !questions || questions.length === 0) return null
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center gap-2">
          Pertanyaan Screening
          <Badge variant="outline">{questions.length} pertanyaan</Badge>
        </CardTitle>
        <CardDescription>
          Pertanyaan khusus untuk proses screening
        </CardDescription>
      </CardHeader>
      <CardContent className="space-y-6">
        {questions.map((q, index) => (
          <div key={q.id}>
            <FormField
              control={form.control}
              name={q.id}
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-base font-medium">
                    {index + 1}. {q.label}
                  </FormLabel>
                  <FormControl>
                    {q.type === "text" || q.type === "date" ? (
                      <Input
                        type={q.type}
                        {...field}
                        value={typeof field.value === 'string' ? field.value : ''}
                        placeholder={`Masukkan ${q.label.toLowerCase()}`}
                      />
                    ) : q.type === "select" ? (
                      <Select
                        onValueChange={field.onChange}
                        value={typeof field.value === 'string' ? field.value : ''}
                      >
                        <SelectTrigger>
                          <SelectValue placeholder="Pilih jawaban" />
                        </SelectTrigger>
                        <SelectContent>
                          {q.options?.map((opt, idx) => (
                            <SelectItem key={idx} value={opt}>
                              {opt}
                            </SelectItem>
                          ))}
                        </SelectContent>
                      </Select>
                    ) : q.type === "checkbox" || q.type === "checkbox_textarea" ? (
                      <div className="space-y-4">
                        <div className="grid grid-cols-1 gap-3">
                          {q.options?.map((opt, idx) => {
                            const currentValues = form.watch(q.id) || []
                            const isChecked = Array.isArray(currentValues) && currentValues.includes(opt)
                            return (
                              <div key={idx} className="flex items-center space-x-2">
                                <Checkbox
                                  id={`${q.id}_${idx}`}
                                  checked={isChecked}
                                  onCheckedChange={(checked) => {
                                    const current = form.getValues(q.id) || []
                                    const currentArray = Array.isArray(current) ? current : []
                                    if (checked) {
                                      form.setValue(q.id, [...currentArray, opt])
                                    } else {
                                      form.setValue(
                                        q.id,
                                        currentArray.filter((v: string) => v !== opt)
                                      )
                                    }
                                  }}
                                />
                                <label
                                  htmlFor={`${q.id}_${idx}`}
                                  className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
                                >
                                  {opt}
                                </label>
                              </div>
                            )
                          })}
                        </div>
                        {/* Textarea untuk "Ya" */}
                        {q.type === "checkbox_textarea" && (
                          (() => {
                            const currentValues = form.watch(q.id) || []
                            const hasYes = Array.isArray(currentValues) &&
                              currentValues.some(val => val.toLowerCase().includes('ya'))
                            return hasYes && (
                              <div className="ml-6 pt-2">
                                <FormLabel className="text-sm">
                                  Mohon jelaskan lebih detail:
                                </FormLabel>
                                <Textarea
                                  placeholder="Tuliskan penjelasan di sini..."
                                  {...form.register(`${q.id}_text`)}
                                  className="mt-2"
                                  rows={3}
                                />
                              </div>
                            )
                          })()
                        )}
                      </div>
                    ) : null}
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            {index < questions.length - 1 && <Separator className="mt-6" />}
          </div>
        ))}
      </CardContent>
    </Card>
  )
} 